import { useEffect, useState } from 'react';

import { JwtToken } from '../api/services/admin/models/JwtToken';

export function useJwtLocalStorage() {
  const [jwtToken, setJwtToken] = useState<JwtToken | null>(null);

  useEffect(() => {
    const jwtLocalStorage = localStorage.getItem('jwt-token');
    if (jwtLocalStorage) {
      setJwtToken(JSON.parse(jwtLocalStorage) as JwtToken);
    }
  }, []);

  return jwtToken;
}
