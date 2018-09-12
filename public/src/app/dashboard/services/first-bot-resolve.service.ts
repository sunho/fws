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
import { PopupService } from '../../core/services/popup.service';
import { AuthService } from '../../core/services/auth.service';

@Injectable({
  providedIn: 'root',
})
export class FirstBotResolverService implements Resolve<Bot> {
  constructor(private botService: BotService,
    private authService: AuthService, private popupService: PopupService, private router: Router) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<Bot> {
    return this.botService.getBots().pipe(
      map(bots => {
        if (bots.length === 0) {
          this.authService.logout().subscribe(
            _ => {
              this.router.navigate(['/']);
            }, error => {
              this.popupService.createMsg(`unknown error(${error}`);
            }
          );
          this.popupService.createMsg('You don\'t own any bot. Contact admin.');
          return null;
        }
        return bots[0];
      })
    );
  }
}
