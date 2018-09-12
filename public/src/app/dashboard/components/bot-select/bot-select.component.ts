import { routeName } from './../../dashboard-routing.module';
import { BotService } from './../../services/bot.service';
import { Component, OnInit } from '@angular/core';
import { Bot } from '../../models/bot';
import { ActivatedRoute, Router } from '@angular/router';
import { PopupService } from '../../../core/services/popup.service';

@Component({
  selector: 'app-bot-select',
  templateUrl: './bot-select.component.html',
})
export class BotSelectComponent implements OnInit {
  current: Bot;
  bots: Bot[];
  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private botService: BotService,
    private popupService: PopupService
  ) {}

  ngOnInit(): void {
    this.current = this.route.snapshot.data.bot;
    this.botService.getBots().subscribe(
      bots => {
        this.bots = bots;
      },
      error => {
        this.popupService.createMsg(`unknown error(${error})`);
      }
    );
  }

  onChange(value: string): void {
    this.router.navigate(['/' + routeName, value]);
  }
}
