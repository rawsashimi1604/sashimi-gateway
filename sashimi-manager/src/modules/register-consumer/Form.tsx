import React, { useEffect, useState } from 'react';
import { AiFillInfoCircle } from 'react-icons/ai';
import { IoIosRemoveCircle } from 'react-icons/io';
import { useNavigate } from 'react-router-dom';
import * as yup from 'yup';

import AdminConsumer from '../../api/services/admin/AdminConsumer';
import AdminService from '../../api/services/admin/AdminService';
import { RegisterConsumerBody } from '../../api/services/admin/body/RegisterConsumerBody';
import { GetAllServicesResponse } from '../../api/services/admin/responses/GetAllServices';
import SelectInput from '../../components/input/SelectInput';
import TextInput from '../../components/input/TextInput';
import ToggleInput from '../../components/input/ToggleInput';
import Subheader from '../../components/typography/Subheader';
import LoadingSpinner from '../../components/utils/LoadingSpinner';
import { delay } from '../../utils/delay';

type FormSubmitState = 'submitting' | 'success' | 'error';

type FormSchema = {
  formUsername: string;
  formEnableJwt: boolean;
  formCredentialName?: string;
  formServices: string[];
};

// Define validation schema using yup
const validationSchema = yup.object().shape({
  formUsername: yup.string().required('Consumer username is required.'),
  formEnableJwt: yup.boolean().required('Choosing whether to enable JWT authentication is required.'),
  formServices: yup.array().required('At least one service is required.').min(1, 'At least one service is required.'),
  formCredentialName: yup.string().required('Credential name is required.')
});

function Form() {
  // Setting up states for the inputs
  const [formData, setFormData] = useState<FormSchema>({
    formUsername: '',
    formEnableJwt: false,
    formServices: []
  });
  const [validationErrors, setValidationErrors] = useState<{
    [key: string]: string;
  }>({});
  const [formState, setFormState] = useState<FormSubmitState | null>(null);
  const [selectedService, setSelectedService] = useState<string>('');
  const [services, setServices] = useState<GetAllServicesResponse | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    loadAllServices();
  }, []);

  useEffect(() => {
    if (!selectedService) return;
    if (formData.formServices.find((service) => service === selectedService)) return;
    handleArrayChange('formServices', [...formData.formServices, selectedService]);
  }, [selectedService]);

  const handleChange = (name: string, value: string) => {
    setFormData((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleArrayChange = (name: string, value: string[]) => {
    setFormData((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleToggleChange = (name: string, value: boolean) => {
    setFormData((prevState) => ({ ...prevState, [name]: value }));
  };

  async function loadAllServices() {
    const res = await AdminService.getAllServices();
    setServices(res.data);
  }

  function getServicesDropdown() {
    if (services) {
      return services.services.map((service) => {
        return service.id + ' - ' + service.name;
      });
    }
    return [];
  }

  function removeService(serviceDetails: string) {
    setFormData((prev) => {
      return {
        ...prev,
        formServices: formData.formServices.filter((svc) => svc !== serviceDetails)
      };
    });
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setFormState('submitting');
    try {
      await validationSchema.validate(formData, { abortEarly: false });
      console.log('Form is valid. Submitting:', formData);
      setValidationErrors({});
    } catch (err) {
      if (err instanceof yup.ValidationError) {
        const errorObj: { [key: string]: string } = {};
        for (let error of err.inner) {
          errorObj[error.path as string] = error.message;
        }
        setValidationErrors(errorObj);
        setFormState(null);
      }
    }
    try {
      await delay(500);
      const body: RegisterConsumerBody = {
        username: formData.formUsername,
        services: formData.formServices.map((val) => parseInt(val.split('-')[0].trim())),
        enableJwtAuth: formData.formEnableJwt,
        jwtCredentialsName: formData.formCredentialName ? formData.formCredentialName : ''
      };
      console.log({ body });
      const res = await AdminConsumer.registerConsumer(body);
      console.log({ res });
      setFormState('success');
      await delay(2000);
      navigate('/consumers');
    } catch (err) {
      setFormState('error');
    }
  };

  return (
    <div className="font-sans">
      <form className="flex flex-col gap-3 w-3/5" onSubmit={handleSubmit}>
        {/* General details */}
        <div className="mt-1">
          <Subheader text="General" align="left" size="sm" />
          <div className="border-b" />
        </div>

        {/* username */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-username" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">username</span>
            <AiFillInfoCircle />
          </label>

          <div className="">
            <TextInput
              id="form-username"
              name="form-username"
              value={formData.formUsername}
              onChange={(e) => handleChange('formUsername', e.target.value)}
              error={validationErrors.formUsername}
            />
          </div>
        </div>

        <div className="mt-1">
          <Subheader text="Configurations" align="left" size="sm" />
          <div className="border-b" />
        </div>

        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-services" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">services</span>
            <AiFillInfoCircle />
          </label>

          <div className="">
            <SelectInput
              options={getServicesDropdown()}
              value={selectedService}
              onChange={(e) => {
                setSelectedService(e);
              }}
              error={validationErrors.formServices}
            />
          </div>
        </div>

        <div className="font-sans text-sm tracking-wider mb-6">
          <div className="flex flex-row items-center gap-3 p-3.5 bg-sashimi-gray/50 rounded-xl">
            {formData.formServices.length > 0 ? (
              formData.formServices.map((svcDetails: string) => {
                return (
                  <span
                    className="animate__animated animate__fadeIn px-2 py-1 rounded-lg bg-sashimi-gray shadow-md flex items-center gap-2 transiton-all duration-150 hover:-translate-y-1 hover:cursor-pointer hover:bg-sashimi-pink"
                    onClick={() => removeService(svcDetails)}
                  >
                    <span>{svcDetails}</span>
                    <IoIosRemoveCircle className="w-5 h-5" />
                  </span>
                );
              })
            ) : (
              <div className="flex items-center justify-center w-full">
                No services registered. You require at least one service.
              </div>
            )}
          </div>
        </div>

        {/* Jwt auth */}
        <div className="flex flex-row items-start justify-between mb-2">
          <div>
            <label htmlFor="form-enableJwt" className="tracking-wide flex flex-row items-center justify-start gap-3">
              <span className="text-sm">enable jwt authentication</span>
            </label>
            <span className="font-sans text-sashimi-deepgray text-xs block">
              enable JWT authentication for all services registered to this consumer.
            </span>
          </div>
          <ToggleInput
            id="form-enableJwt"
            name="form-enableJwt"
            checked={formData.formEnableJwt}
            onChange={(e) => handleToggleChange('formEnableJwt', e)}
          />
        </div>

        {formData.formEnableJwt && (
          <div className="flex flex-col justify-center gap-1 text-sm">
            <label
              htmlFor="form-credentialname"
              className="tracking-wide flex flex-row items-center justify-start gap-3"
            >
              <span className="mb-1">credential name</span>
              <AiFillInfoCircle />
            </label>

            <div className="">
              <TextInput
                id="form-credentialname"
                name="form-credentialname"
                value={formData.formCredentialName}
                onChange={(e) => handleChange('formCredentialName', e.target.value)}
                error={validationErrors.formCredentialName}
              />
            </div>
          </div>
        )}

        <button
          type="submit"
          className="w-[80px] mt-2 text-xs py-1.5 px-2 pb-2 text-white bg-sashimi-deepgreen shadow-md rounded-lg font-sans tracking-wider border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg"
        >
          <span>register</span>
        </button>

        {formState == 'submitting' && (
          <div className="flex flex-row items-center gap-2 text-sm tracking-wider">
            <React.Fragment>
              <span className="text-sashimi-deepyellow">registering your consumer...</span>
              <LoadingSpinner size={12} />
            </React.Fragment>
          </div>
        )}

        {formState == 'success' && (
          <div className="flex flex-row items-center gap-2 text-sm tracking-wider">
            <React.Fragment>
              <span className="text-sashimi-deepgreen">consumer registration success! redirecting...</span>
              <LoadingSpinner size={12} />
            </React.Fragment>
          </div>
        )}

        {formState == 'error' && (
          <div className="flex flex-row items-center gap-2 text-sm tracking-wider">
            <React.Fragment>
              <span className="text-sashimi-deeppink">failed to register consumer. please try again.</span>
            </React.Fragment>
          </div>
        )}
      </form>
    </div>
  );
}

export default Form;
