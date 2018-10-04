import { Volume } from './../models/bot';
import { Injectable } from '@angular/core';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';
import { BotService } from './bot.service';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class VolumeResolverService implements Resolve<Volume> {
  constructor(private botService: BotService, private router: Router) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<Volume> {
    const bot = route.parent.data['bot'];
    return this.botService.getVolumes(bot.id).pipe(
      map(vols => {
        const vol = vols.find(v => v.name === route.params.name);
        if (!vol) {
          this.router.navigate(['..']);
        }
        return vol;
      })
    );
  }
}
