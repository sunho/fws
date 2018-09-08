import { NgModule } from '@angular/core';
import { RouterModule, Routes, PreloadAllModules } from '@angular/router';
import { LoginComponent } from './core/pages/login/login.component';
import { InviteComponent } from './core/pages/invite/invite.component';

const routes: Routes = [
  {path: '', redirectTo: '/', pathMatch: 'full'},
  {path: '', component: LoginComponent},
  {path: 'invite', component: InviteComponent},
  {path: 'dashboard', loadChildren: './dashboard/dashboard.module#DashBoardModule'}
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes,
      {
        scrollPositionRestoration: 'enabled'
      })
  ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }
