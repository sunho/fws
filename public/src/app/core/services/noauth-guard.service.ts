import { map, catchError } from "rxjs/operators"
import { Observable } from "rxjs"
import { AuthService } from "./auth.service"
import { CanActivate, Router } from "@angular/router"
import { Injectable } from "@angular/core"
import { routeName } from "../../dashboard/dashboard-routing.module"

@Injectable({
  providedIn: "root",
})
export class NoAuthGuardService implements CanActivate {
  constructor(private authService: AuthService, private router: Router) {}

  canActivate(): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.authService.getUser().subscribe(
        _ => {
          observer.next(false)
          observer.complete()
          this.router.navigate(["/" + routeName])
        },
        _ => {
          observer.next(true)
          observer.complete()
        }
      )
    })
  }
}
