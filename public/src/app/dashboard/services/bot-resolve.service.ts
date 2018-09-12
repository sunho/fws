import { environment } from './../../../environments/environment';
import { BotService } from './bot.service';
import { Injectable } from '@angular/core';
import {
  Resolve,
  ActivatedRouteSnapshot,
  RouterStateSnapshot,
  Router,
} from '@angular/router';
import { Bot } from '../models/bot';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class BotResolverService implements Resolve<Bot> {
  constructor(private botService: BotService, private router: Router) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<Bot> {
    return this.botService.getBots().pipe(
      map(bots => {
        const bot = bots.find(b => b.id === parseInt(route.params.id, 10));
        if (!bot) {
          this.router.navigate(['/' + environment.dashboardRoute]);
        }
        return bot;
      })
    );
  }
}
