import { BotResolverService as BotResolver } from './services/bot-resolve.service';
import { EnvResolverService as EnvResolver } from './services/env-resolve.service';
import { FirstBotRedirectService as FirstBotRedirect } from './services/first-bot-resolve.service';
import { Routes, RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { DashComponent } from './pages/dash/dash.component';
import { HomeComponent } from './pages/home/home.component';
import { ConfigComponent } from './pages/config/config.component';
import { VolumeComponent } from './pages/volume/volume.component';
import { EnvDetailComponent } from './pages/env-detail/env-detail.component';

const routes: Routes = [
  {
    path: '',
    component: DashComponent,
    canActivate: [FirstBotRedirect],
  },
  {
    path: ':id',
    component: DashComponent,
    resolve: {
      bot: BotResolver,
    },
    children: [
      { path: '', redirectTo: 'home', pathMatch: 'full' },
      { path: 'home', component: HomeComponent },
      { path: 'volume', component: VolumeComponent },
      { path: 'config', component: ConfigComponent },
      {
        path: 'env/:name',
        component: EnvDetailComponent,
        resolve: {
          env: EnvResolver,
        },
      },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class DashBoardRoutingModule {}
