import { Component, OnInit } from '@angular/core';
import { MatDialogRef, MatDialogConfig, MatDialog } from '@angular/material/dialog';
import { App } from 'src/app/models/app';
import { AppService } from 'src/app/services/app.service';
import { BroadcasterService } from 'src/app/services/broadcaster.service';
import { PasswordDialogComponent } from 'src/app/password-dialog/password-dialog.component';

@Component({
  selector: 'app-app-dialog',
  templateUrl: './app-dialog.component.html',
  styleUrls: ['./app-dialog.component.scss']
})
export class AppDialogComponent implements OnInit {

  apps: App[] = [];

  constructor(
    public dialogRef: MatDialogRef<AppDialogComponent>,
    private service: AppService,
    private broadcasterService: BroadcasterService,
    private ddialog: MatDialog) { }

  ngOnInit(): void {
    this.service.getAllApps('default').subscribe(apps => {
      this.apps = this.service.sort(apps);
      if (this.apps == null) {
        return;
      }
      this.apps.forEach(el => {
        el.inprogress = this.service.installing.get(el.name);
      });
    });
  }

  install(app: App): void {
    const dialog = new MatDialogConfig();
    dialog.autoFocus = true;
    dialog.disableClose = true;
    dialog.data = {};
    const dialogRef = this.ddialog.open(PasswordDialogComponent, dialog);

    dialogRef.afterClosed().subscribe(
      data => {
        if (data !== undefined) {
          this.dialogRef.close();
          this.service.installing.set(app.name, true);
          app.user = data.user;
          app.passwd = data.pass;
          this.service.installApp('default', app).subscribe(() => {
            this.broadcasterService.broadcast('reload-apps');
            this.service.installing.set(app.name, false);
          }, () => {
            this.broadcasterService.broadcast('reload-apps');
            this.service.installing.set(app.name, false);
          });

        }
      });
  }
}
