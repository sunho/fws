import { ActivatedRoute } from '@angular/router';
import { BotService } from './../../services/bot.service';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Bot } from '../../models/bot';

@Component({
    selector: 'app-header',
    templateUrl: './header.component.html'
})
export class HeaderComponent implements OnInit {
    current: Bot;

    constructor(private botService: BotService, private route: ActivatedRoute) { }

    ngOnInit(): void {
        this.route.data.subscribe(d => {
            this.current = d.bot;
        }, _ => { });
    }

    onRebuildClick(): void {
        // TODO
        console.log(this.current.id);
        this.botService.rebuildBot(this.current.id).subscribe(
            _ => {},
            err => {}
        );
    }
}
