import React, { useState } from 'react';
import { AiFillInfoCircle } from 'react-icons/ai';
import { useNavigate } from 'react-router-dom';
import * as yup from 'yup';

import AdminService from '../../api/services/admin/AdminService';
import TextAreaInput from '../../components/input/TextAreaInput';
import TextInput from '../../components/input/TextInput';
import ToggleInput from '../../components/input/ToggleInput';
import Subheader from '../../components/typography/Subheader';
import LoadingSpinner from '../../components/utils/LoadingSpinner';
import { delay } from '../../utils/delay';

type FormSubmitState = 'submitting' | 'success' | 'error';

const isValidUrl = (value: string) => {
  try {
    const url = new URL(value);
    return ['http:', 'https:'].includes(url.protocol);
  } catch (_) {
    return false;
  }
};

// Define validation schema using yup
const validationSchema = yup.object().shape({
  formName: yup.string().required('Service name is required.'),
  formTargetUrl: yup
    .string()
    .required('Target URL is required.')
    .test('is-valid-url', 'Must be a valid URL.', isValidUrl),
  formPath: yup.string().required('Gateway path is required.'),
  formDescription: yup.string().required('Service description is required.')
});

function Form() {
  // Setting up states for the inputs
  const [formData, setFormData] = useState({
    formName: '',
    formTargetUrl: '',
    formPath: '',
    formDescription: '',
    formHealthChecks: false
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
      const registerRes = await AdminService.registerService({
        name: formData.formName,
        targetUrl: formData.formTargetUrl,
        path: formData.formPath,
        description: formData.formDescription,
        healthCheckEnabled: formData.formHealthChecks
      });
      console.log(registerRes);
      setFormState('success');
      await delay(2000);
      navigate('/services');
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
        {/* Service name */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-name" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">service name</span>
            <AiFillInfoCircle />
          </label>
          <div className="">
            <TextInput
              id="form-name"
              name="form-name"
              value={formData.formName}
              onChange={(e) => handleChange('formName', e.target.value)}
              error={validationErrors.formName}
            />
          </div>
        </div>

        {/* TargetUrl */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-name" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">target url</span>
            <AiFillInfoCircle />
          </label>
          <div className="">
            <TextInput
              id="form-targetUrl"
              name="form-targetUrl"
              value={formData.formTargetUrl}
              onChange={(e) => handleChange('formTargetUrl', e.target.value)}
              error={validationErrors.formTargetUrl}
            />
          </div>
        </div>

        {/* Path */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-path" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">gateway path</span>
            <AiFillInfoCircle />
          </label>
          <div className="">
            <TextInput
              id="form-path"
              name="form-path"
              value={formData.formPath}
              onChange={(e) => handleChange('formPath', e.target.value)}
              error={validationErrors.formPath}
            />
          </div>
        </div>

        {/* Description */}
        <div className="flex flex-col justify-center gap-1 text-sm">
          <label htmlFor="form-description" className="tracking-wide flex flex-row items-center justify-start gap-3">
            <span className="mb-1">service description</span>
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
            <label htmlFor="form-description" className="tracking-wide flex flex-row items-center justify-start gap-3">
              <span className="text-sm">use swagger open api specification</span>
            </label>
            <span className="font-sans text-sashimi-deepgray text-xs block">
              some boiler plate text here. to be implemented
            </span>
          </div>
          <ToggleInput
            id="form-swagger"
            name="form-swagger"
            checked={false}
            onChange={(e) => console.log('to be implemented.')}
            disabled
          />
        </div>

        {/* Health checks */}
        <div className="flex flex-row items-start justify-between mb-2">
          <div>
            <label htmlFor="form-description" className="tracking-wide flex flex-row items-center justify-start gap-3">
              <span className="text-sm">enable health checks</span>
            </label>
            <span className="font-sans text-sashimi-deepgray text-xs block">
              To enable health checks, service must contain a <span className="italic">'/healthz'</span> endpoint that
              returns a 200 OK status code.
            </span>
          </div>
          <ToggleInput
            id="form-healthchecks"
            name="form-healthchecks"
            checked={formData.formHealthChecks}
            onChange={(e) => handleToggleChange('formHealthChecks', e)}
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
              <span className="text-sashimi-deepyellow">registering your service...</span>
              <LoadingSpinner size={12} />
            </React.Fragment>
          </div>
        )}

        {formState == 'success' && (
          <div className="flex flex-row items-center gap-2 text-sm tracking-wider">
            <React.Fragment>
              <span className="text-sashimi-deepgreen">service registration success! redirecting...</span>
              <LoadingSpinner size={12} />
            </React.Fragment>
          </div>
        )}

        {formState == 'error' && (
          <div className="flex flex-row items-center gap-2 text-sm tracking-wider">
            <React.Fragment>
              <span className="text-sashimi-deeppink">failed to register service. please try again.</span>
            </React.Fragment>
          </div>
        )}
      </form>
    </div>
  );
}

export default Form;
