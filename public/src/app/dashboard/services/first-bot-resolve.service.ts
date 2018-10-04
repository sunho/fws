import { BotService } from './bot.service';
import { Injectable } from '@angular/core';
import { Router, CanActivate } from '@angular/router';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { PopupService } from '../../core/services/popup.service';
import { AuthService } from '../../core/services/auth.service';
import { environment } from '../../../environments/environment.prod';

@Injectable({
  providedIn: 'root',
})
export class FirstBotRedirectService implements CanActivate {
  constructor(
    private botService: BotService,
    private authService: AuthService,
    private popupService: PopupService,
    private router: Router
  ) {}

  canActivate(): Observable<boolean> {
    return this.botService.getBots().pipe(
      map(bots => {
        if (bots.length === 0) {
          this.authService.logout().subscribe(
            _ => {
              this.popupService.createMsg(
                'You don\'t own any bot. Contact admin.'
              );
              this.router.navigate(['/']);
            },
            error => {
              this.popupService.createMsg(`unknown error(${error}`);
            }
          );
          return false;
        }
        this.router.navigate(['/' + environment.dashboardRoute, bots[0].id]);
        return false;
      })
    );
  }
}
