import { ActivatedRoute } from '@angular/router';
import { ModalService } from './../../services/modal.service';
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

  envOptions = [{ title: 'Edit', func: this.envEditCallback.bind(this) }, { title: 'Delete', func: this.envDeleteCallback.bind(this) }];

  envItems: Env[];

  confNames = ['name', 'path', 'file'];

  confKeys = ['name', 'path', 'file'];

  confOptions = [{ title: 'Edit', func: this.confEditCallback.bind(this) }, { title: 'Delete', func: this.confDeleteCallback.bind(this) }];

  confItems: Config[];

  constructor(
    private botService: BotService,
    private popupService: PopupService,
    private modalService: ModalService,
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

  confEditCallback(conf: Config, s: string): void {
    this.modalService.createMod({
      title: 'Edit Config File',
      items: [
        {
          name: 'name',
          key: 'name',
          initial: conf.name,
          disabled: true
        },
        {
          name: 'path',
          key: 'path',
          initial: conf.path
        },
        {
          name: 'file',
          key: 'file',
          initial: conf.file
        },
        {
          name: 'value',
          key: 'value',
          initial: conf.value,
          textfield: true
        },
      ],
      button: 'Edit',
      callback: this.confEditCompleteCallback.bind(this),
    });
  }


  envEditCallback(env: Env, s: string): void {
    this.modalService.createMod({
      title: 'Edit Environment Variable',
      items: [
        {
          name: 'name',
          key: 'name',
          initial: env.name,
          disabled: true
        },
        {
          name: 'value',
          key: 'value',
          initial: env.value
        },
      ],
      button: 'Edit',
      callback: this.envEditCompleteCallback.bind(this),
    });
  }

  confEditCompleteCallback(obj: Object): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.botService.patchConfig(this.current.id, obj as Config).subscribe(
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

  envEditCompleteCallback(obj: Object): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.botService.patchEnv(this.current.id, obj as Env).subscribe(
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

  onConfAddClick(): boolean {
    this.modalService.createMod({
      title: 'Add Config File',
      items: [
        {
          name: 'name',
          key: 'name'
        },
        {
          name: 'path',
          key: 'path'
        },
        {
          name: 'file',
          key: 'file'
        },
        {
          name: 'value',
          key: 'value',
          textfield: true
        }
      ],
      callback: this.confAddCallback.bind(this),
      button: 'Add',
    });
    return false;
  }

  onEnvAddClick(): boolean {
    this.modalService.createMod({

      title: 'Add Environment Variable',
      items: [
        {
          name: 'name',
          key: 'name'
        },
        {
          name: 'value',
          key: 'value'
        }
      ],
      callback: this.envAddCallback.bind(this),
      button: 'Add',
    });
    return false;
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
