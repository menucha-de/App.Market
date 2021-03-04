import { Component, OnInit, OnDestroy, ViewChild } from '@angular/core';
import { Subscription } from 'rxjs';
import { AppService } from '../services/app.service';
import { App } from '../models/app';
import { BroadcasterService } from '../services/broadcaster.service';
import { MatMenuTrigger } from '@angular/material/menu';
import { ToastService } from '../toast/toast.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import { MatDialogConfig, MatDialog } from '@angular/material/dialog';
import { PasswordDialogComponent } from 'src/app/password-dialog/password-dialog.component';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss']
})
export class MainComponent implements OnInit, OnDestroy {
  private reloadApps: Subscription;
  private loading: Subscription;
  isLoading = false;
  apps$: App[] = [];
  @ViewChild('xx', { read: MatMenuTrigger })
  contextMenu: MatMenuTrigger;
  contextMenuPosition = { x: '0px', y: '0px' };

  constructor(
    private service: AppService,
    private broadcasterService: BroadcasterService,
    private toast: ToastService,
    private snackBar: MatSnackBar,
    private ddialog: MatDialog) { }

  ngOnInit(): void {
    this.service.getInstalledApps('default').subscribe(apps => {
      this.apps$ = this.service.sort(apps);
    });

    this.reloadApps = this.broadcasterService.on('reload-apps').subscribe(() => {
      this.service.getInstalledApps('default').subscribe(apps => {
        this.apps$ = this.service.sort(apps)
      });
    });

    this.loading = this.broadcasterService.on('loading').subscribe((data: boolean) => {
      this.isLoading = data;
    });
  }

  ngOnDestroy() {
    if (this.reloadApps) {
      this.reloadApps.unsubscribe();
    }
    if (this.loading) {
      this.loading.unsubscribe();
    }
  }

  onContextMenu(event: any, app: App) {
    const el = document.getElementById(app.label);
    event.preventDefault();
    this.contextMenuPosition.x = el.offsetLeft + 'px';
    this.contextMenuPosition.y = (el.offsetTop + 95) + 'px';
    this.contextMenu.menuData = { app };
    this.contextMenu.menu.focusFirstItem('mouse');
    this.contextMenu.openMenu();
  }

  onContextMenuDelete(app: App) {
    this.service.deleteApp('default', app.name).subscribe(() => {
      this.broadcasterService.broadcast('reload-apps');
    }, err => {
      this.toast.openSnackBar(err.error, 'Error');
    });
  }

  onContextMenuStart(app: App) {
    app.state = 'STARTING';
    this.service.updateApp('default', app.name, app).subscribe(() => {
      this.broadcasterService.broadcast('reload-apps');
    }, err => {
      this.toast.openSnackBar(err.error, 'Error');
    });
  }

  onContextMenuStop(app: App) {
    app.state = 'STOPPING';
    this.service.updateApp('default', app.name, app).subscribe(() => {
      this.broadcasterService.broadcast('reload-apps');
    }, err => {
      this.toast.openSnackBar(err.error, 'Error');
    });
  }

  onContextMenuUpgrade(app: App) {
    const dialog = new MatDialogConfig();
    dialog.autoFocus = true;
    dialog.data = {};
    const dialogRef = this.ddialog.open(PasswordDialogComponent, dialog);

    dialogRef.afterClosed().subscribe(
      data => {
        if (data !== undefined) {
          app.user = data.user;
          app.passwd = data.pass;
          if (app.state == 'STARTED') {
            app.state = 'UPGRADINGSTARTED';
          } else {
            app.state = 'UPGRADINGSTOPPED';
          }
          this.service.installing.set(app.name, true);
          this.service.updateApp('default', app.name, app).subscribe(() => {
            this.broadcasterService.broadcast('reload-apps');
            this.service.installing.set(app.name, false);
          }, () => {
            this.broadcasterService.broadcast('reload-apps');
            this.service.installing.set(app.name, false);
          });
        }
      }
    );
  }

  onContextMenuReset(app: App) {
    if (app.state == 'STARTED') {
      app.state = 'RESETTINGSTARTED';
    } else {
      app.state = 'RESETTINGSTOPPED';
    }
    this.service.installing.set(app.name, true);
    this.service.updateApp('default', app.name, app).subscribe(() => {
      this.broadcasterService.broadcast('reload-apps');
      this.service.installing.set(app.name, false);
    }, () => {
      this.broadcasterService.broadcast('reload-apps');
      this.service.installing.set(app.name, false);
    });
  }

  showApp(name: string) {
    location.href = `/${name}/`;
  }
}
