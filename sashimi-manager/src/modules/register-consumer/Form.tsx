import React, { useEffect, useState } from 'react';
import { AiFillInfoCircle } from 'react-icons/ai';
import { useNavigate } from 'react-router-dom';
import * as yup from 'yup';

import AdminConsumer from '../../api/services/admin/AdminConsumer';
import AdminRoute from '../../api/services/admin/AdminRoute';
import AdminService from '../../api/services/admin/AdminService';
import { RegisterConsumerBody } from '../../api/services/admin/body/RegisterConsumerBody';
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
  formUsername: yup.string().required('Consumer username is required.')
});

function Form() {
  // Setting up states for the inputs
  const [formData, setFormData] = useState({
    formUsername: ''
  });
  const [validationErrors, setValidationErrors] = useState<{
    [key: string]: string;
  }>({});
  const [formState, setFormState] = useState<FormSubmitState | null>(null);
  const navigate = useNavigate();

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
      const body: RegisterConsumerBody = {
        username: formData.formUsername
      };
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

        {/* Service */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-username" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">consumer username</span>
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

        <h2 className="tracking-wide text-sm">No consumer configurations.</h2>

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
