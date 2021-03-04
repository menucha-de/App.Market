import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-mat-spinner-overlay',
  templateUrl: './mat-spinner-overlay.component.html',
  styleUrls: ['./mat-spinner-overlay.component.scss']
})
export class MatSpinnerOverlayComponent implements OnInit {

  constructor() { }
  @Input() loading = false;
  ngOnInit(): void {

  }

}
