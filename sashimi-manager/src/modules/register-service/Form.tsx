import React, { useState } from 'react';
import { AiFillInfoCircle } from 'react-icons/ai';
import * as yup from 'yup';

import TextInput from '../../components/input/TextInput';

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
    formDescription: ''
  });
  const [validationErrors, setValidationErrors] = useState<{
    [key: string]: string;
  }>({});
  const [apiRequestError, setApiRequestError] = useState('');

  const handleChange = (name: string, value: string) => {
    setFormData((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      await validationSchema.validate(formData, { abortEarly: false });
      console.log('Form is valid. Submitting:', formData);
      setValidationErrors({});
      setApiRequestError('');
      // Submit the form
    } catch (err) {
      if (err instanceof yup.ValidationError) {
        const errorObj: { [key: string]: string } = {};
        for (let error of err.inner) {
          errorObj[error.path as string] = error.message;
        }
        setValidationErrors(errorObj);
      }
    }
  };

  return (
    <div className="font-sans">
      <form className="flex flex-col gap-3 w-3/5" onSubmit={handleSubmit}>
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
              // Add error prop to display error message, assuming your TextInput component supports it
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
              // Add error prop to display error message, assuming your TextInput component supports it
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
              // Add error prop to display error message, assuming your TextInput component supports it
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
            <TextInput
              id="form-description"
              name="form-description"
              value={formData.formDescription}
              onChange={(e) => handleChange('formDescription', e.target.value)}
              // Add error prop to display error message, assuming your TextInput component supports it
              error={validationErrors.formDescription}
            />
          </div>
        </div>

        <button
          type="submit"
          className="w-[80px] mt-2 text-xs py-1.5 px-2 pb-2 text-white bg-sashimi-deepgreen shadow-md rounded-lg font-sans tracking-wider border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg"
        >
          <span>register</span>
        </button>
      </form>
    </div>
  );
}

export default Form;
