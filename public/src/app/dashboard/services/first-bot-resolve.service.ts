import { BotService } from './bot.service';
import { Injectable } from '@angular/core';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { Bot } from '../models/bot';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class FirstBotResolverService implements Resolve<Bot> {
  constructor(private botService: BotService) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<Bot> {
    return this.botService.getBots().pipe(map(bots => {
      if (bots.length === 0) {
        throw new Error('no bot');
      }
      return bots[0];
    }));
  }
}
