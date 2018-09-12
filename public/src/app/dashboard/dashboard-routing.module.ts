import { BotResolverService as BotResolver } from './services/bot-resolve.service';
import { FirstBotRedirectService as FirstBotRedirect } from './services/first-bot-resolve.service';
import { Routes, RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { BuildComponent } from './pages/build/build.component';
import { DashComponent } from './pages/dash/dash.component';
import { HomeComponent } from './pages/home/home.component';
import { ConfigComponent } from './pages/config/config.component';
import { VolumeComponent } from './pages/volume/volume.component';

const routes: Routes = [
  {
    path: '',
    component: DashComponent,
    canActivate: [
      FirstBotRedirect,
    ]
  },
  {
    path: ':id',
    component: DashComponent,
    resolve: {
      bot: BotResolver,
    },
    children: [
      { path: '', redirectTo: 'home', pathMatch: 'full'},
      { path: 'home', component: HomeComponent },
      { path: 'build', component: BuildComponent },
      { path: 'volume', component: VolumeComponent },
      { path: 'config', component: ConfigComponent },
    ]
  },
];

export const routeName = 'dashboard';

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class DashBoardRoutingModule {}
