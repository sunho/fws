import { PopupService } from './../../../core/services/popup.service';
import { ActivatedRoute } from '@angular/router';
import { BotService, CONFLICT } from './../../services/bot.service';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Bot } from '../../models/bot';
import { STRINGS } from '../../../../locale/strings';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
})
export class HeaderComponent implements OnInit {
  current: Bot;
  bots: Bot[] = [];

  constructor(
    private botService: BotService,
    private route: ActivatedRoute,
    private popupService: PopupService
  ) {}

  ngOnInit(): void {
    this.route.data.subscribe(
      d => {
        this.current = d.bot;
      },
      _ => {}
    );

    this.botService.getBots().subscribe(
      bots => {
        this.bots = bots;
      },
      error => {
        this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
      }
    );
  }

  onRebuildClick(): boolean {
    this.botService.rebuildBot(this.current.id).subscribe(
      _ => {},
      error => {
        this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
      }
    );
    return false;
  }

  onUploadClick(): boolean {
    this.botService.uploadBot(this.current.id).subscribe(
      _ => {},
      error => {
        this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
      }
    );
    return false;
  }

  onRestartClick(): boolean {
    this.botService.restartBot(this.current.id).subscribe(
      _ => {},
      error => {
        this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
      }
    );
    return false;
  }
}
