import { Component, OnInit } from '@angular/core';
import { Subscription } from 'rxjs/internal/Subscription';
import { ActivatedRoute } from '@angular/router';
import { Env } from '../../models/bot';

@Component({
  selector: 'app-env-detail',
  templateUrl: './env-detail.component.html',
})
export class EnvDetailComponent implements OnInit {
  subscription: Subscription;

  current: Env;

  constructor(
    private route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    this.subscription = this.route.data.subscribe(d => {
      this.current = d['env'];
    });
  }
}
