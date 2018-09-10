import { NgModule } from '@angular/core';
import { RouterModule, Routes, PreloadAllModules } from '@angular/router';
import { LoginComponent } from './core/pages/login/login.component';
import { InviteComponent } from './core/pages/invite/invite.component';
import { routeName as dashboardRouteName } from './dashboard/dashboard-routing.module';
import { AuthGuardService as AuthGuard } from './core/services/auth-guard.service';

const routes: Routes = [
  {path: '', redirectTo: '/', pathMatch: 'full'},
  {path: '', component: LoginComponent},
  {path: 'invite', component: InviteComponent},
  {path: dashboardRouteName, canActivate: [AuthGuard], loadChildren: './dashboard/dashboard.module#DashBoardModule'}
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
