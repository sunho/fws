import { PopupService } from './../../../core/services/popup.service';
import { map, catchError, startWith } from 'rxjs/operators';
import { ActivatedRoute } from '@angular/router';
import { Component, OnInit, OnDestroy, ViewEncapsulation } from '@angular/core';
import { Bot, BuildStatus, RunStatus } from '../../models/bot';
import { Observable, Observer, Subscription, interval } from 'rxjs';
import { BotService, NOT_FOUND } from '../../services/bot.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
})
export class HomeComponent implements OnInit, OnDestroy {
  current: Bot;
  buildStatus: BuildStatus = { type: 'none' };
  runStatus: RunStatus = { type: 'none' };
  subscription: Subscription;
  subscription2: Subscription;
  constructor(
    private botService: BotService,
    private route: ActivatedRoute,
    private popupService: PopupService
  ) {}

  ngOnInit(): void {
    this.subscription2 = this.route.parent.data.subscribe(d => {
      this.current = d.bot;
      if (this.subscription) {
        this.subscription.unsubscribe();
      }
      this.subscription = interval(1000)
        .pipe(startWith(0))
        .subscribe(_ => {
          this.botService.getBuildStatus(this.current.id).subscribe(
            s => {
              this.buildStatus = s;
            },
            error => {
              if (error === NOT_FOUND) {
                this.buildStatus = { type: 'none' };
              }
            }
          );
          this.botService.getRunStatus(this.current.id).subscribe(
            s => {
              this.runStatus = s;
            },
            error => {
              if (error === NOT_FOUND) {
                this.runStatus = { type: 'none' };
              }
            }
          );
        });
    });
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
    this.subscription2.unsubscribe();
  }
}
