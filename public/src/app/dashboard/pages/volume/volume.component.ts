import { Observable } from 'rxjs';
import { AddModalService } from './../../services/add-modal.service';
import { PopupService } from './../../../core/services/popup.service';
import { BotService, CONFLICT, BAD_FORMAT } from './../../services/bot.service';
import { Component, OnInit } from '@angular/core';
import { Volume, Bot } from '../../models/bot';
import { ActivatedRoute } from '@angular/router';
import { STRINGS } from '../../../../locale/strings';

@Component({
  selector: 'app-volume',
  templateUrl: './volume.component.html',
})
export class VolumeComponent implements OnInit {
  constructor(
    private route: ActivatedRoute,
    private botService: BotService,
    private popupService: PopupService,
    private addModalService: AddModalService
  ) {}

  current: Bot;

  resNames = ['name', 'path'];

  resKeys = ['name', 'path'];

  resOptions = [{ title: 'Delete', func: this.deleteCallback.bind(this) }];

  resItems: Volume[];

  deleteCallback(vol: Volume, s: string): void {
    this.botService.deleteVolume(this.current.id, vol.name).subscribe(
      _ => {
        this.refreshItems();
      },
      error => {
        this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
      }
    );
  }

  addCallback(obj: object): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.botService.addVolume(this.current.id, obj as Volume).subscribe(
        _ => {
          this.refreshItems();
          observer.next(true);
          observer.complete();
        },
        error => {
          if (error === CONFLICT) {
            this.popupService.createMsg(STRINGS.EXIST_VOLUME);
          } else if (error === BAD_FORMAT) {
            this.popupService.createMsg(STRINGS.BAD_VOLUME_FORMAT);
          } else {
            this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
          }
        }
      );
    });
  }

  onAddClick(): boolean {
    this.addModalService.createMod({
      title: 'Add Volume',
      keys: ['name', 'path'],
      names: ['name', 'path'],
      callback: this.addCallback.bind(this),
    });
    return false;
  }

  refreshItems(): void {
    this.botService.getVolumes(this.current.id).subscribe(
      vols => {
        this.resItems = vols;
      },
      error => {
        this.popupService.createMsg(`unknown error ${error}`);
      }
    );
  }

  ngOnInit(): void {
    this.route.parent.data.subscribe(d => {
      this.current = d.bot;
      this.refreshItems();
    });
  }
}
