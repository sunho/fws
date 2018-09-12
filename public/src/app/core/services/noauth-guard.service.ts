import { environment } from './../../../environments/environment';
import { map, catchError } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { AuthService } from './auth.service';
import { CanActivate, Router } from '@angular/router';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class NoAuthGuardService implements CanActivate {
  constructor(private authService: AuthService, private router: Router) {}

  canActivate(): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.authService.getUser().subscribe(
        _ => {
          observer.next(false);
          observer.complete();
          this.router.navigate(['/' + environment.dashboardRoute]);
        },
        _ => {
          observer.next(true);
          observer.complete();
        }
      );
    });
  }
}
