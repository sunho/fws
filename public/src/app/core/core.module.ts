import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { LoginComponent } from './pages/login/login.component';
import { LoginFormComponent } from './components/login-form/login-form.component';
import { InviteComponent } from './pages/invite/invite.component';
import { InviteFormComponent } from './components/invite-form/invite-form.component';
import { SmallHeaderComponent } from './components/small-header/small-header.component';

@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [
    LoginComponent,
    LoginFormComponent,
    InviteComponent,
    InviteFormComponent,
    SmallHeaderComponent
  ],
  exports: [
    LoginComponent,
    InviteComponent
  ]
})
export class CoreModule { }
