import { ServiceHealth } from '../types/api/ServiceHealth.interface';

export function serviceHealthColor(health: ServiceHealth) {
  console.log({ health });
  switch (health) {
    case 'healthy':
      return {
        bg: 'bg-sashimi-deepgreen',
        text: 'text-sashimi-deepgreen'
      };
    case 'unhealthy':
      return {
        bg: 'bg-sashimi-deeppink',
        text: 'text-sashimi-deeppink'
      };
    case 'startup':
      return {
        bg: 'bg-sashimi-deepyellow',
        text: 'text-sashimi-deepyellow'
      };
    case 'not_enabled':
      return {
        bg: 'bg-sashimi-deepgray',
        text: 'text-sashimi-deepgray'
      };
    default:
      return {
        bg: 'bg-sashimi-deepgray',
        text: 'text-sashimi-deepgray'
      };
  }
}
