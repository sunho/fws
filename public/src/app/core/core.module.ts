import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { HttpModule } from '@angular/http';

import { FormInputComponent } from './components/form-input/form-input.component';
import { LoginComponent } from './pages/login/login.component';
import { LoginFormComponent } from './components/login-form/login-form.component';
import { InviteComponent } from './pages/invite/invite.component';
import { InviteFormComponent } from './components/invite-form/invite-form.component';
import { SmallHeaderComponent } from './components/small-header/small-header.component';
import { IconComponent } from './components/icon/icon.component';
import { AuthGuardService } from './services/auth-guard.service';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    HttpClientModule,
    HttpModule,
    FormsModule
  ],
  declarations: [
    FormInputComponent,
    IconComponent,
    LoginComponent,
    LoginFormComponent,
    InviteComponent,
    InviteFormComponent,
    SmallHeaderComponent
  ],
  providers: [
  ],
  exports: [
    LoginComponent,
    InviteComponent,
    FormInputComponent,
    IconComponent
  ]
})
export class CoreModule { }
