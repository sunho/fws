import {InjectionToken} from '@angular/core';

interface IAppConfig {
  siteName: string;
  apiUrl: string;
}

export let APP_CONFIG = new InjectionToken('app.config');

export const AppConfig: IAppConfig = {
  siteName: 'SWS',
  apiUrl: 'http://localhost:8080/api',
};

