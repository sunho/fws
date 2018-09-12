import { AuthService } from "./auth.service"
import { Injectable } from "@angular/core"
import { CanActivate, Router } from "@angular/router"
import { Observable, throwError } from "rxjs"
import { map, catchError } from "rxjs/operators"

@Injectable({
  providedIn: "root",
})
export class AuthGuardService implements CanActivate {
  constructor(private authService: AuthService, private router: Router) {}

  canActivate(): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.authService.getUser().subscribe(
        _ => {
          observer.next(true)
          observer.complete()
        },
        _ => {
          observer.next(false)
          observer.complete()
          this.router.navigate(["/"])
        }
      )
    })
  }
}
