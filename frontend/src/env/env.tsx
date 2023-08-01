import development from './env.development';
import production  from './env.production';

export const env = process.env.NODE_ENV === 'production' ? production : development;
