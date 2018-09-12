import { PopupService } from './../../../core/services/popup.service';
import { BotService } from './../../services/bot.service';
import { Component, OnInit } from '@angular/core';
import { Volume, Bot } from '../../models/bot';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-volume',
  templateUrl: './volume.component.html',
})
export class VolumeComponent implements OnInit {
  current: Bot;

  resNames = [
    'name',
    'path',
  ];

  resKeys = [
    'name',
    'path',
  ];

  resItems: Volume[];

  constructor(private route: ActivatedRoute, private botService: BotService, private popupService: PopupService) {}

  ngOnInit(): void {
    this.route.parent.data.subscribe(d => {
      this.current = d.bot;
      this.botService.getVolumes(this.current.id).subscribe(
        vols => {
          this.resItems = vols;
        },
        error => {
          this.popupService.createMsg(`unknown error ${error}`);
        }
      );
    });
  }
}
