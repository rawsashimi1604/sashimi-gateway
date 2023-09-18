import React, { useEffect, useState } from 'react';
import { AiFillInfoCircle } from 'react-icons/ai';
import { useNavigate } from 'react-router-dom';
import * as yup from 'yup';

import AdminRoute from '../../api/services/admin/AdminRoute';
import AdminService from '../../api/services/admin/AdminService';
import { RegisterRouteBody } from '../../api/services/admin/body/RegisterRouteBody';
import { GetAllServicesResponse } from '../../api/services/admin/responses/GetAllServices';
import SelectInput from '../../components/input/SelectInput';
import TextAreaInput from '../../components/input/TextAreaInput';
import TextInput from '../../components/input/TextInput';
import ToggleInput from '../../components/input/ToggleInput';
import Subheader from '../../components/typography/Subheader';
import LoadingSpinner from '../../components/utils/LoadingSpinner';
import { delay } from '../../utils/delay';

type FormSubmitState = 'submitting' | 'success' | 'error';

// Define validation schema using yup
const validationSchema = yup.object().shape({
  formServiceId: yup.string().required('Service id is required.'),
  formPath: yup.string().required('Gateway route path is required.'),
  formDescription: yup.string().required('Route description is required.'),
  formMethod: yup.string().required('Route method is required.')
});

function Form() {
  // Setting up states for the inputs
  const [formData, setFormData] = useState({
    formServiceId: '',
    formPath: '',
    formDescription: '',
    formMethod: ''
  });
  const [validationErrors, setValidationErrors] = useState<{
    [key: string]: string;
  }>({});
  const [formState, setFormState] = useState<FormSubmitState | null>(null);
  const [services, setServices] = useState<GetAllServicesResponse | null>(null);
  const navigate = useNavigate();

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

  const handleChange = (name: string, value: string) => {
    setFormData((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleToggleChange = (name: string, value: boolean) => {
    setFormData((prevState) => ({ ...prevState, [name]: value }));
  };

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
      const body: RegisterRouteBody = {
        serviceId: Number(formData.formServiceId.split('-')[0].trim()),
        path: formData.formPath,
        description: formData.formDescription,
        method: formData.formMethod
      };
      const res = await AdminRoute.registerRoute(body);
      console.log({ res });
      setFormState('success');
      await delay(2000);
      navigate('/routes');
    } catch (err) {
      setFormState('error');
    }
  };

  useEffect(() => {
    loadAllServices();
  }, []);

  return (
    <div className="font-sans">
      <form className="flex flex-col gap-3 w-3/5" onSubmit={handleSubmit}>
        {/* General details */}
        <div className="mt-1">
          <Subheader text="General" align="left" size="sm" />
          <div className="border-b" />
        </div>

        {/* Service */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-serviceId" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">service</span>
            <AiFillInfoCircle />
          </label>

          <div className="">
            <SelectInput
              options={getServicesDropdown()}
              onChange={(e) => handleChange('formServiceId', e)}
              value={formData.formServiceId}
              error={validationErrors.formServiceId}
            />
          </div>
        </div>

        {/* Route path */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-routePath" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">route path</span>
            <AiFillInfoCircle />
          </label>

          <div className="">
            <TextInput
              id="form-routePath"
              name="form-routePath"
              value={formData.formPath}
              onChange={(e) => handleChange('formPath', e.target.value)}
              error={validationErrors.formPath}
            />
          </div>
        </div>

        {/* Route method */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-method" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">method</span>
            <AiFillInfoCircle />
          </label>

          <div className="">
            <SelectInput
              options={['GET', 'POST', 'PATCH', 'DELETE', 'OPTIONS']}
              onChange={(e) => handleChange('formMethod', e)}
              value={formData.formMethod}
              error={validationErrors.formMethod}
            />
          </div>
        </div>

        {/* Route description */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-description" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">route description</span>
            <AiFillInfoCircle />
          </label>

          <div className="">
            <TextAreaInput
              id="form-description"
              name="form-description"
              value={formData.formDescription}
              onChange={(e) => handleChange('formDescription', e.target.value)}
              error={validationErrors.formDescription}
            />
          </div>
        </div>

        <div className="mt-1">
          <Subheader text="Configurations" align="left" size="sm" />
          <div className="border-b" />
        </div>

        {/* Health checks */}
        <div className="flex flex-row items-start justify-between mb-2">
          <div>
            <label htmlFor="form-auth" className="tracking-wide flex flex-row items-center justify-start gap-3">
              <span className="text-sm">enable JWT authentication</span>
            </label>
            <span className="font-sans text-sashimi-deepgray text-xs block">
              Enable JWT Authentication for route (WIP)
            </span>
          </div>
          <ToggleInput
            id="form-auth"
            name="form-auth"
            checked={true}
            disabled
            onChange={(e) => handleToggleChange('formAuth', e)}
          />
        </div>

        <button
          type="submit"
          className="w-[80px] mt-2 text-xs py-1.5 px-2 pb-2 text-white bg-sashimi-deepgreen shadow-md rounded-lg font-sans tracking-wider border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg"
        >
          <span>register</span>
        </button>

        {formState == 'submitting' && (
          <div className="flex flex-row items-center gap-2 text-sm tracking-wider">
            <React.Fragment>
              <span className="text-sashimi-deepyellow">registering your route...</span>
              <LoadingSpinner size={12} />
            </React.Fragment>
          </div>
        )}

        {formState == 'success' && (
          <div className="flex flex-row items-center gap-2 text-sm tracking-wider">
            <React.Fragment>
              <span className="text-sashimi-deepgreen">route registration success! redirecting...</span>
              <LoadingSpinner size={12} />
            </React.Fragment>
          </div>
        )}

        {formState == 'error' && (
          <div className="flex flex-row items-center gap-2 text-sm tracking-wider">
            <React.Fragment>
              <span className="text-sashimi-deeppink">failed to register route. please try again.</span>
            </React.Fragment>
          </div>
        )}
      </form>
    </div>
  );
}

export default Form;
