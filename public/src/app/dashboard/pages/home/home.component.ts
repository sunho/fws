import { PopupService } from "./../../../core/services/popup.service"
import { map, catchError, startWith } from "rxjs/operators"
import { ActivatedRoute } from "@angular/router"
import { Component, OnInit, OnDestroy } from "@angular/core"
import { Bot, BuildStatus } from "../../models/bot"
import { Observable, Observer, Subscription, interval } from "rxjs"
import { BotService, NOT_FOUND } from "../../services/bot.service"

@Component({
  selector: "app-home",
  templateUrl: "./home.component.html",
})
export class HomeComponent implements OnInit, OnDestroy {
  current: Bot
  buildStatus: BuildStatus = { type: "none" }
  subscription: Subscription
  constructor(
    private botService: BotService,
    private route: ActivatedRoute,
    private popupService: PopupService
  ) {
    this.route.data.subscribe(d => {
      this.current = d.bot
      if (this.subscription) {
        this.subscription.unsubscribe()
      }
      this.subscription = interval(500)
        .pipe(startWith(0))
        .subscribe(_ => {
          this.botService.getBotBuildStatus(this.current.id).subscribe(
            s => {
              this.buildStatus = s
            },
            error => {
              if (error === NOT_FOUND) {
                this.buildStatus = { type: "none" }
              }
            }
          )
        })
    })
  }

  ngOnInit(): void {}

  ngOnDestroy(): void {
    this.subscription.unsubscribe()
  }
}
