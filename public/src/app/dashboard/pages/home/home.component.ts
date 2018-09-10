import { map, catchError } from 'rxjs/operators';
import { ActivatedRoute } from '@angular/router';
import { Component, OnInit } from '@angular/core';
import { Bot, BuildStatus } from '../../models/bot';
import { Observable, Observer } from 'rxjs';
import { BotService } from '../../services/bot.service';

@Component({
    selector: 'app-home',
    templateUrl: './home.component.html'
})
export class HomeComponent implements OnInit {
    $current: Observable<Bot>;
    $buildStatus: Observable<BuildStatus>;
    $buildStatusObserver: Observer<BuildStatus>;
    constructor(private botService: BotService, private route: ActivatedRoute) {
        this.$current = this.route.data.pipe(map(d => d.bot));
        this.$buildStatus = new Observable<BuildStatus>(observer => {
            this.$buildStatusObserver = observer;
        });
        this.$current.subscribe(b => {
            this.botService.getBotBuildStatus(b.id).
                subscribe(s => {
                        this.$buildStatusObserver.next(s);
                    }, err => {
                        this.$buildStatusObserver.next({type: 'none'});
                    }
                );
        });
    }

    ngOnInit(): void { }
}
