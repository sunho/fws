import { routeName as dashboardRouteName } from './../../../dashboard/dashboard-routing.module';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html'
})
export class LoginComponent implements OnInit {
  constructor(private router: Router, private authSerivce: AuthService) {
  }

  ngOnInit() {
    this.authSerivce.getUser().subscribe(
      _ => {
        this.router.navigate([dashboardRouteName]);
      },
      _ => { });
  }

}
