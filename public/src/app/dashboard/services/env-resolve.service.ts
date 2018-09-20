import { environment } from './../../../environments/environment';
import { BotService } from './bot.service';
import { Injectable } from '@angular/core';
import {
  Resolve,
  ActivatedRouteSnapshot,
  RouterStateSnapshot,
  Router,
} from '@angular/router';
import { Bot, Env } from '../models/bot';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class EnvResolverService implements Resolve<Env> {
  constructor(private botService: BotService, private router: Router) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<Env> {
    return this.botService.getEnvs(route.parent.data['bot'].id).pipe(
      map(envs => {
        const env = envs.find(e => e.name === route.params.name);
        if (!env) {
          this.router.navigate(['..']);
        }
        return env;
      })
    );
  }
}
