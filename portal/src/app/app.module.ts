import {BrowserModule} from '@angular/platform-browser';
import {LOCALE_ID, NgModule} from '@angular/core';

import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {CommonModule} from '@angular/common';
import {registerLocaleData} from '@angular/common';
import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {HttpClientModule} from '@angular/common/http';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {NZ_I18N} from 'ng-zorro-antd/i18n';
import {zh_CN} from 'ng-zorro-antd/i18n';
import zh from '@angular/common/locales/zh';
import {LoginComponent} from './login/login.component';
import {IconsProviderModule} from './icons-provider.module';
import {PageNotFoundComponent} from './page-not-found/page-not-found.component';
import {NzInputModule} from 'ng-zorro-antd/input';
import {NzButtonModule} from 'ng-zorro-antd/button';
import {NzFormModule} from 'ng-zorro-antd/form';
import {NzCheckboxModule} from 'ng-zorro-antd/checkbox';
import {NzIconModule} from 'ng-zorro-antd/icon';
import {NzMessageModule} from 'ng-zorro-antd/message';
import {NzListModule} from 'ng-zorro-antd/list';
import {NzModalModule} from 'ng-zorro-antd/modal';
import {NgxAmapModule} from "ngx-amap";
import {InstallComponent} from "./install/install.component";
import {NzStepsModule} from "ng-zorro-antd/steps";
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {MarkdownModule} from "ngx-markdown";
import {NzSelectModule} from "ng-zorro-antd/select";
import {NzTableModule} from "ng-zorro-antd/table";
import {HelperModule} from "./helper/helper.module";
import {NZ_CONFIG, NzConfig} from "ng-zorro-antd/core/config";
//import {HelperModule} from "./helper/helper.module";

registerLocaleData(zh, 'zh');

export function tokenGetter() {
  return localStorage.getItem('token');
}

const ngZorroConfig: NzConfig = {
  //不要关闭
  modal: {nzMaskClosable: false}
};


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    InstallComponent,
    PageNotFoundComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    NzButtonModule,
    NzFormModule,
    NzInputModule,
    NzCheckboxModule,
    NzIconModule,
    NzMessageModule,
    NzModalModule,
    NzListModule,
    IconsProviderModule,
    HttpClientModule,
    NgxAmapModule.forRoot({apiKey: 'e4c1bd11fe1b25d77dae4cf3993f7034', debug: true}),
    NzStepsModule,
    NzLayoutModule,
    MarkdownModule.forRoot(),
    NzSelectModule,
    NzTableModule,
    HelperModule,
    //HelperModule,
  ],
  providers: [
    {provide: NZ_CONFIG, useValue: ngZorroConfig},
    {provide: NZ_I18N, useValue: zh_CN},
    {provide: LOCALE_ID, useValue: 'zh'}],
  bootstrap: [AppComponent]
})
export class AppModule {
}
