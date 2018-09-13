import { ActivatedRoute } from '@angular/router';
import { AddModalService } from './../../services/add-modal.service';
import { PopupService } from './../../../core/services/popup.service';
import { BotService, CONFLICT } from './../../services/bot.service';
import { Component, OnInit } from '@angular/core';
import { Config, Env, Bot } from '../../models/bot';
import { Observable } from 'rxjs';
import { STRINGS } from '../../../../locale/strings';

@Component({
  selector: 'app-config',
  templateUrl: './config.component.html',
})
export class ConfigComponent implements OnInit {
  current: Bot;

  envNames = ['name', 'value'];

  envKeys = ['name', 'value'];

  envOptions = [{ title: 'Delete', func: this.envDeleteCallback.bind(this) }];

  envItems: Env[];

  confNames = ['name', 'path', 'file'];

  confKeys = ['name', 'path', 'file'];

  confOptions = [{ title: 'Delete', func: this.confDeleteCallback.bind(this) }];

  confItems: Config[];

  constructor(
    private botService: BotService,
    private popupService: PopupService,
    private addModalService: AddModalService,
    private route: ActivatedRoute
  ) {}

  confAddCallback(obj: object): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.botService.addConfig(this.current.id, obj as Config).subscribe(
        _ => {
          this.refreshItems();
          observer.next(true);
          observer.complete();
        },
        error => {
          if (error === CONFLICT) {
            this.popupService.createMsg(STRINGS.EXIST_CONFIG);
          } else {
            this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
          }
        }
      );
    });
  }

  envAddCallback(obj: object): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.botService.addEnv(this.current.id, obj as Env).subscribe(
        _ => {
          this.refreshItems();
          observer.next(true);
          observer.complete();
        },
        error => {
          if (error === CONFLICT) {
            this.popupService.createMsg(STRINGS.EXIST_ENV);
          } else {
            this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
          }
        }
      );
    });
  }

  confDeleteCallback(conf: Config, s: string): void {
    this.botService.deleteConfig(this.current.id, conf.name).subscribe(
      _ => {
        this.refreshItems();
      },
      error => {
        this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
      }
    );
  }

  envDeleteCallback(env: Env, s: string): void {
    this.botService.deleteEnv(this.current.id, env.name).subscribe(
      _ => {
        this.refreshItems();
      },
      error => {
        this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
      }
    );
  }

  onConfAddClick(): void {
    this.addModalService.createMod({
      title: 'Add Config File',
      keys: ['name', 'path', 'file'],
      names: ['name', 'path', 'file'],
      callback: this.confAddCallback.bind(this),
    });
  }

  onEnvAddClick(): void {
    this.addModalService.createMod({
      title: 'Add Environment Variable',
      keys: ['name', 'value'],
      names: ['name', 'value'],
      callback: this.envAddCallback.bind(this),
    });
  }

  ngOnInit(): void {
    this.route.parent.data.subscribe(d => {
      this.current = d.bot;
      this.refreshItems();
    });
  }

  refreshItems(): void {
    this.botService.getEnvs(this.current.id).subscribe(
      envs => {
        this.envItems = envs;
      },
      error => {
        this.popupService.createMsg(`unknown error ${error}`);
      }
    );

    this.botService.getConfigs(this.current.id).subscribe(
      confs => {
        this.confItems = confs;
      },
      error => {
        this.popupService.createMsg(`unknown error ${error}`);
      }
    );
  }
}
