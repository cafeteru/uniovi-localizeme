import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ModalStageComponent } from '../modal-stage/modal-stage.component';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { Stage } from '../../../types/stage';
import { sortStagesByActive, sortStagesByName } from '../../../shared/sorts/stages-sorts';
import { StageService } from '../../../core/services/stage.service';
import { tap } from 'rxjs';
import { NzModalService } from 'ng-zorro-antd/modal';

@Component({
    selector: 'app-stage-list',
    templateUrl: './stage-list.component.html',
    styleUrls: ['./stage-list.component.scss'],
})
export class StageListComponent extends BaseComponent implements OnInit {
    currentPageStages: readonly Stage[] = [];
    stages: readonly Stage[] = [];
    isLoading = false;

    listOfColumns: ColumnHeader<Stage>[] = [
        {
            name: 'Name',
            sortOrder: null,
            sortFn: sortStagesByName,
            sortDirections,
        },
        {
            name: 'Active',
            sortOrder: null,
            sortFn: sortStagesByActive,
            sortDirections,
        },
    ];

    constructor(
        private nzMessageService: NzMessageService,
        private nzModalService: NzModalService,
        private stageService: StageService,
        public matDialog: MatDialog
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.loadStages();
    }

    loadStages(): void {
        this.stages = [];
        this.isLoading = true;
        const subscription$ = this.stageService.findAll().subscribe({
            next: (stages) => (this.stages = stages),
            error: () => (this.isLoading = false),
            complete: () => (this.isLoading = false),
        });
        this.subscriptions$.push(subscription$);
    }

    onCurrentPageDataChange($event: readonly Stage[]): void {
        this.currentPageStages = $event;
    }

    openModal(stage?: Stage): void {
        const newStage: Stage = {
            name: undefined,
            active: true,
            id: undefined,
        };
        const dialogRef = this.matDialog.open(ModalStageComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: stage ? stage : newStage,
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result: Stage) => {
            if (result) {
                this.loadStages();
            }
        });
        this.subscriptions$.push(subscription$);
    }

    disable(stage: Stage): void {
        const subscription$ = this.stageService.disable(stage).subscribe({
            next: () => this.loadStages(),
            error: () => this.nzMessageService.create('error', 'Error disabling'),
        });
        this.subscriptions$.push(subscription$);
    }

    showDeleteModal(stage: Stage): void {
        this.nzModalService.confirm({
            nzTitle: 'Are you sure delete this stage?',
            nzOkText: 'Yes',
            nzOkType: 'primary',
            nzOkDanger: true,
            nzOnOk: () => this.delete(stage),
            nzCancelText: 'No',
            nzAutofocus: 'cancel',
        });
    }

    private delete(stage: Stage): void {
        const subscription$ = this.stageService.delete(stage).subscribe((result) => {
            if (result) {
                this.loadStages();
                this.nzMessageService.create('success', `${stage.name} has been deleted`);
            } else {
                this.nzMessageService.create('error', 'Error deleting');
            }
        });
        this.subscriptions$.push(subscription$);
    }
}
