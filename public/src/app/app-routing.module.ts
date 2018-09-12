import { AppConfig } from './app.config';
import { NgModule } from '@angular/core';
import { RouterModule, Routes, PreloadAllModules } from '@angular/router';
import { LoginComponent } from './core/pages/login/login.component';
import { InviteComponent } from './core/pages/invite/invite.component';
import { AuthGuardService as AuthGuard } from './core/services/auth-guard.service';
import { NoAuthGuardService as NoAuthGuard } from './core/services/noauth-guard.service';

const routes: Routes = [
  { path: '', redirectTo: '/', pathMatch: 'full' },
  { path: '', canActivate: [NoAuthGuard], component: LoginComponent },
  { path: 'invite', canActivate: [NoAuthGuard], component: InviteComponent },
  {
    path: AppConfig.dashboardRoute,
    canActivate: [AuthGuard],
    loadChildren: './dashboard/dashboard.module#DashBoardModule',
  },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {
      scrollPositionRestoration: 'enabled',
    }),
  ],
  exports: [RouterModule],
})
export class AppRoutingModule {}
