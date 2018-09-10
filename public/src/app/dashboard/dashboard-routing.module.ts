import { BotResolverService as BotResolver } from './services/bot-resolve.service';
import { FirstBotResolverService as FirstBotResolver } from './services/first-bot-resolve.service';
import { Routes, RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { BuildComponent } from './pages/build/build.component';
import { DashComponent } from './pages/dash/dash.component';
import { HomeComponent } from './pages/home/home.component';

const routeChildren = [
  { path: '', component: HomeComponent },
  { path: 'build', component: BuildComponent },
];

const routes: Routes = [
  {
    path: '',
    component: DashComponent,
    resolve: {
      bot: FirstBotResolver,
    },
    children: routeChildren,
  },
  {
    path: ':id',
    component: DashComponent,
    resolve: {
      bot: BotResolver,
    },
    children: routeChildren,
  },
];

export const routeName = 'dashboard';

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class DashBoardRoutingModule {}
