import {NgModule} from '@angular/core';

import {IconsProviderModule} from './icons-provider.module';
import {NzLayoutModule} from 'ng-zorro-antd/layout';
import {NzMenuModule} from 'ng-zorro-antd/menu';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {CommonModule} from '@angular/common';
import {HttpClientModule} from '@angular/common/http';
import {AdminRoutingModule} from './admin-routing.module';

import {NzSpaceModule} from 'ng-zorro-antd/space';
import {NzIconModule} from 'ng-zorro-antd/icon';
import {NzToolTipModule} from 'ng-zorro-antd/tooltip';
import {NzTableModule} from 'ng-zorro-antd/table';
import {NzModalModule} from 'ng-zorro-antd/modal';
import {NzButtonModule} from 'ng-zorro-antd/button';
import {NzCheckboxModule} from 'ng-zorro-antd/checkbox';
import {NzSwitchModule} from 'ng-zorro-antd/switch';
import {NzPopconfirmModule} from 'ng-zorro-antd/popconfirm';
import {NzDividerModule} from 'ng-zorro-antd/divider';
import {NzDrawerModule} from 'ng-zorro-antd/drawer';
import {NzSelectModule} from 'ng-zorro-antd/select';
import {NzInputNumberModule} from 'ng-zorro-antd/input-number';
import {NzStatisticModule} from 'ng-zorro-antd/statistic';
import {NzCollapseModule} from 'ng-zorro-antd/collapse';
import {NzFormModule} from 'ng-zorro-antd/form';
import {NzInputModule} from 'ng-zorro-antd/input';
import {NzTabsModule} from 'ng-zorro-antd/tabs';
import {NzTransferModule} from 'ng-zorro-antd/transfer';
import {NzRadioModule} from 'ng-zorro-antd/radio';
import {NzProgressModule} from 'ng-zorro-antd/progress';
import {NzCardModule} from 'ng-zorro-antd/card';
import {NzUploadModule} from 'ng-zorro-antd/upload';
import {NzDropDownModule} from "ng-zorro-antd/dropdown";
import {NzDatePickerModule} from "ng-zorro-antd/date-picker";
import {NzTimePickerModule} from "ng-zorro-antd/time-picker";
import {DragDropModule} from "@angular/cdk/drag-drop";
import {NgxEchartsModule} from "ngx-echarts";
import {NzGridModule} from "ng-zorro-antd/grid";
import {NgxAmapModule} from "ngx-amap";
// import {NgxEchartsModule} from 'ngx-echarts';
//import * as echarts from 'echarts';

import {AdminComponent} from "./admin.component";

import {WelcomeComponent} from "./welcome/welcome.component";
import {UnknownComponent} from "./unknown/unknown.component";
import {DashComponent} from "./dash/dash.component";
import {HomeComponent} from "./home/home.component";
import {TunnelComponent} from "./tunnel/tunnel.component";
import {LinkComponent} from "./link/link.component";
import {DeviceComponent} from "./device/device.component";
import {ElementComponent} from "./element/element.component";
import {ProjectComponent} from "./project/project.component";
import {TemplateComponent} from "./template/template.component";
import {PluginComponent} from "./plugin/plugin.component";
import {ProtocolComponent} from "./protocol/protocol.component";
import {SettingComponent} from "./setting/setting.component";
import {UserComponent} from "./user/user.component";
import {PasswordComponent} from "./password/password.component";
import {TunnelDetailComponent} from "./tunnel-detail/tunnel-detail.component";
import {TunnelEditComponent} from "./tunnel-edit/tunnel-edit.component";
import {LinkDetailComponent} from "./link-detail/link-detail.component";
import {LinkEditComponent} from "./link-edit/link-edit.component";
import {DeviceDetailComponent} from "./device-detail/device-detail.component";
import {DeviceEditComponent} from "./device-edit/device-edit.component";
import {ElementDetailComponent} from "./element-detail/element-detail.component";
import {ElementEditComponent} from "./element-edit/element-edit.component";
import {ProjectDetailComponent} from "./project-detail/project-detail.component";
import {ProjectEditComponent} from "./project-edit/project-edit.component";
import {TemplateDetailComponent} from "./template-detail/template-detail.component";
import {TemplateEditComponent} from "./template-edit/template-edit.component";
import {EditJobsComponent} from "./edit-jobs/edit-jobs.component";
import {EditPollersComponent} from "./edit-pollers/edit-pollers.component";
import {EditMappingComponent} from "./edit-mapping/edit-mapping.component";
import {EditCalculatorsComponent} from "./edit-calculators/edit-calculators.component";
import {EditCommandsComponent} from "./edit-commands/edit-commands.component";
import {EditAggregatorsComponent} from "./edit-aggregators/edit-aggregators.component";
import {EditStrategiesComponent} from "./edit-strategies/edit-strategies.component";
import {HelperModule} from "../helper/helper.module";
import {NzBreadCrumbModule} from "ng-zorro-antd/breadcrumb";
import {ContainerComponent} from "./container/container.component";
import {EditRegisterComponent} from "./edit-register/edit-register.component";
import {EditHeartbeatComponent} from "./edit-heartbeat/edit-heartbeat.component";
import {EditProtocolComponent} from "./edit-protocol/edit-protocol.component";
import {TunnelEditDevicesComponent} from "./tunnel-edit-devices/tunnel-edit-devices.component";

@NgModule({
  declarations: [
    AdminComponent,
    ContainerComponent,
    WelcomeComponent,
    UnknownComponent,
    DashComponent,
    HomeComponent,
    TunnelComponent, TunnelDetailComponent, TunnelEditComponent,
    EditRegisterComponent, EditHeartbeatComponent, EditProtocolComponent, TunnelEditDevicesComponent,
    LinkComponent, LinkDetailComponent, LinkEditComponent,
    DeviceComponent, DeviceDetailComponent, DeviceEditComponent,
    ElementComponent, ElementDetailComponent, ElementEditComponent,
    ProjectComponent, ProjectDetailComponent, ProjectEditComponent,
    TemplateComponent, TemplateDetailComponent, TemplateEditComponent,
    EditMappingComponent, EditPollersComponent, EditJobsComponent, EditStrategiesComponent,
    EditCalculatorsComponent, EditCommandsComponent, EditAggregatorsComponent,
    PluginComponent,
    ProtocolComponent,
    SettingComponent,
    UserComponent,
    PasswordComponent,
  ],
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        HttpClientModule,
        AdminRoutingModule,
        IconsProviderModule,
        NzIconModule,
        NzGridModule,
        NzLayoutModule,
        NzMenuModule,
        NzToolTipModule,
        NzTableModule,
        NzModalModule,
        NzFormModule,
        NzButtonModule,
        NzInputModule,
        NzCheckboxModule,
        NzSwitchModule,
        NzPopconfirmModule,
        NzDividerModule,
        NzDrawerModule,
        NzSelectModule,
        NzSpaceModule,
        NzInputNumberModule,
        NzStatisticModule,
        NzTabsModule,
        NzCollapseModule,
        NzTransferModule,
        NzRadioModule,
        NzProgressModule,
        NzCardModule,
        NzUploadModule,
        NzDropDownModule,
        NzTimePickerModule,
        NzDatePickerModule,
        DragDropModule,

        NgxEchartsModule.forRoot({echarts: () => import('echarts')}),

        NgxAmapModule.forRoot({apiKey: 'e4c1bd11fe1b25d77dae4cf3993f7034', debug: true}),
        HelperModule,
        NzBreadCrumbModule,
    ],
  bootstrap: [AdminComponent],
  providers: []
})
export class AdminModule {
}
