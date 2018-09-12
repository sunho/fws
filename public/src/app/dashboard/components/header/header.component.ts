import { PopupService } from './../../../core/services/popup.service';
import { ActivatedRoute } from '@angular/router';
import { BotService, CONFLICT } from './../../services/bot.service';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Bot } from '../../models/bot';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
})
export class HeaderComponent implements OnInit {
  current: Bot;

  constructor(private botService: BotService, private route: ActivatedRoute, private popupService: PopupService) {}

  ngOnInit(): void {
    this.route.data.subscribe(
      d => {
        this.current = d.bot;
      },
      _ => {}
    );
  }

  onRebuildClick(): void {
    this.botService.rebuildBot(this.current.id).subscribe(_ => {}, error => {
      if (error === CONFLICT) {
        this.popupService.createMsg('a build is already in progress');
      }
    });
  }
}
