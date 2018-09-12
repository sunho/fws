import { environment } from './../../../../environments/environment';
import { BotService } from './../../services/bot.service';
import { Component, OnInit, Input } from '@angular/core';
import { Bot } from '../../models/bot';
import { ActivatedRoute, Router } from '@angular/router';
import { PopupService } from '../../../core/services/popup.service';

@Component({
  selector: 'app-bot-select',
  templateUrl: './bot-select.component.html',
})
export class BotSelectComponent implements OnInit {
  @Input()
  current: number;
  @Input()
  items: Bot[];
  constructor(
    private router: Router,
  ) {}

  ngOnInit(): void {
  }

  onChange(value: string): void {
    const r = this.router.url.replace(`${environment.dashboardRoute}/${this.current}`, `${environment.dashboardRoute}/${value}`);
    this.router.navigate([r]);
  }
}
