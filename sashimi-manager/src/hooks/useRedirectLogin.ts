import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

// When the user does not have a jwt token, we should redirect to login page
export function useRedirectLogin() {
  const navigate = useNavigate();

  useEffect(() => {
    const jwtToken = localStorage.getItem('jwt-token');
    if (!jwtToken) {
      navigate('/login');
    }
  }, []);
}
