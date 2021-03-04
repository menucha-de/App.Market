import { Component, OnInit,  ElementRef, ViewChild } from '@angular/core';
import { MatDialog, DialogPosition, MatDialogConfig } from '@angular/material/dialog';
import { NameDialogComponent } from '../name-dialog/name-dialog.component';
import { AppService } from '../services/app.service';
import { BroadcasterService } from '../services/broadcaster.service';
import { AppDialogComponent } from './app-dialog/app-dialog.component';

@Component({
  selector: 'app-install',
  templateUrl: './install.component.html',
  styleUrls: ['./install.component.scss']
})
export class InstallComponent implements OnInit {

  @ViewChild('FileSelectInputDialog') FileSelectInputDialog: ElementRef;

  constructor(
    public ddialog: MatDialog,
    private broadcasterService: BroadcasterService,
    private service: AppService) { }

  ngOnInit(): void {
  }

  openAppDialog() {
    this.ddialog.open(AppDialogComponent, {
      width: 'fit-content', position: { top: '10em' } as DialogPosition
    });

  }
  openAddFilesDialog() {
    const e: HTMLElement = this.FileSelectInputDialog.nativeElement;
    e.click();
  }
  onFileAdded(files: FileList) {
    const dialog = new MatDialogConfig();
    dialog.autoFocus = true;
    dialog.disableClose = true;
    dialog.data = {};
    const dialogRef = this.ddialog.open(NameDialogComponent, dialog);
    dialogRef.afterClosed().subscribe(
      data => {
        if (data !== undefined) {
          dialogRef.close();
          this.broadcasterService.broadcast('loading',  true);
          this.service.uploadImgFile('default', files.item(0), data.name).subscribe(() => {
            this.broadcasterService.broadcast('reload-apps');
            this.broadcasterService.broadcast('loading',  false);
          }, () => {
            this.broadcasterService.broadcast('reload-apps');
            this.broadcasterService.broadcast('loading',  false);
          });
          this.FileSelectInputDialog.nativeElement.value = null;
        }
      }
    );
  }
}
