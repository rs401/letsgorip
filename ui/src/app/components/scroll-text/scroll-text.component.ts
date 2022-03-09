import { animate, state, style, transition, trigger } from '@angular/animations';
import { Component, Input, OnInit, SimpleChanges } from '@angular/core';

@Component({
  selector: 'app-scroll-text',
  templateUrl: './scroll-text.component.html',
  styleUrls: ['./scroll-text.component.css'],
  animations: [
    trigger('scrollTrigger', [
      state('bottom', style({'opacity': '0'})),
      state('top', style({'opacity': '1'})),
      transition('void => *', animate('1000ms ease-in-out')),
      transition('bottom => top', animate('1000ms ease-in-out')),
      transition('top => bottom', animate('1000ms ease-in-out')),
    ])
  ]
})
export class ScrollTextComponent implements OnInit {
  // @Input() msg!: string;
  @Input() lines!: string[];
  public msg: string = "";
  public move: boolean = false;

  constructor() {

  }

  ngOnInit(): void {
    let rand = Math.floor(Math.random() * this.lines.length);
    this.msg = this.lines[rand];
    setInterval(() => {
      this.scrollLine();
    }, 3000);
  }

  scrollLine(): void {
    let rand = Math.floor(Math.random() * this.lines.length);
    setTimeout(() => {
      this.move = !this.move;
    }, 2000);
    this.move = !this.move;
    this.msg = this.lines[rand];
  }

}
