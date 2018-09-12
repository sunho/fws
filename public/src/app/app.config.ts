import { InjectionToken } from '@angular/core';

interface IAppConfig {
  siteName: string;
  apiUrl: string;
  dashboardRoute: string;
}

export let APP_CONFIG = new InjectionToken('app.config');

export const AppConfig: IAppConfig = {
  siteName: 'FWS',
  apiUrl: '/api',
  dashboardRoute: 'dashboard',
};
