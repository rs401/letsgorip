import { Component, OnInit, Input, SimpleChanges } from '@angular/core';

@Component({
  selector: 'app-flash-message',
  templateUrl: './flash-message.component.html',
  styleUrls: ['./flash-message.component.css']
})
export class FlashMessageComponent implements OnInit {

  @Input() public message?: string;
  showMessage: boolean;

  constructor() {
    this.showMessage = false;
  }

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    this.manageMessage(changes['message'].currentValue);
}
manageMessage(message: string) {
    if (message && message.length > 0) {
        this.message = message;
        this.showMessage = true;

        setTimeout(() => {
            this.showMessage = false;
            this.message = '';
        }, 3000);
    }
}

}
