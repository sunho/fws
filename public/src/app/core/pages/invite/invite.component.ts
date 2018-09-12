import { Component, OnInit } from "@angular/core"
import { ActivatedRoute, Router } from "@angular/router"

import { AuthService } from "../../services/auth.service"

@Component({
  selector: "app-invite",
  templateUrl: "./invite.component.html",
  styleUrls: ["./invite.component.scss"],
})
export class InviteComponent implements OnInit {
  key: string
  username: string

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private authSerivce: AuthService
  ) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      this.key = params["key"]
      this.username = params["username"]
      this.authSerivce.keyCheck(this.key, this.username).subscribe(
        _ => {},
        error => {
          this.router.navigate(["/"])
        }
      )
    })
  }

  onSuccess(): void {
    this.router.navigate(["/"])
  }
}
