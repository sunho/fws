import { routeName } from './../../dashboard-routing.module';
import { BotService } from './../../services/bot.service';
import { Component, OnInit } from '@angular/core';
import { Bot } from '../../models/bot';
import { ActivatedRoute, Router } from '@angular/router';

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
    private botService: BotService
  ) {}

  ngOnInit(): void {
    this.current = this.route.snapshot.data.bot;
    this.botService.getBots().subscribe(
      bots => {
        this.bots = bots;
      },
      err => {
        console.error(err);
      }
    );
  }

  onChange(value: string): void {
    this.router.navigate(['/' + routeName, value]);
  }
}
