import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CoreModule } from './core/core.module';

import localeKo from '@angular/common/locales/ko';

import { registerLocaleData } from '@angular/common';
import { APP_BASE_HREF } from '@angular/common';

registerLocaleData(localeKo, 'ko');

@NgModule({
  declarations: [AppComponent],
  imports: [BrowserModule, CoreModule, AppRoutingModule],
  providers: [ {provide: APP_BASE_HREF, useValue: 'http://localhost:8080'}],
  bootstrap: [AppComponent],
})
export class AppModule {}
