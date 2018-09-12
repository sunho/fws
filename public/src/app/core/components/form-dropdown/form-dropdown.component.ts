import { Component, OnInit, Input, ViewChild, HostListener, ElementRef } from '@angular/core';

export interface DropdownItem {
    title: string;
    func: (string) => void;
}

@Component({
    selector: 'app-form-dropdown',
    templateUrl: './form-dropdown.component.html'
})
export class FormDropdownComponent implements OnInit {
    @Input() items: DropdownItem[];
    show: boolean;

    onItemClick(e: MouseEvent, item: DropdownItem): void {
        e.preventDefault();
        item.func(item.title);
    }

    onButtonClick(e: MouseEvent): void {
        e.stopPropagation();
        this.show = !this.show;
    }

    @HostListener('document:click', ['$event'])
    onClick(e: MouseEvent): void {
        this.show = false;
    }

    constructor() { }

    ngOnInit(): void { }

}
