// Copyright (c) 2019 the Octant contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//

import {
  ComponentFixture,
  fakeAsync,
  TestBed,
  waitForAsync,
} from '@angular/core/testing';

import {
  HeptagonGridRowComponent,
  HoverStatus,
} from './heptagon-grid-row.component';
import { HeptagonComponent } from '../../smart/heptagon/heptagon.component';
import { windowProvider, WindowToken } from '../../../../../window';
import { SharedModule } from '../../../shared.module';

describe('HeptagonGridRowComponent', () => {
  let component: HeptagonGridRowComponent;
  let fixture: ComponentFixture<HeptagonGridRowComponent>;

  beforeEach(
    waitForAsync(() => {
      TestBed.configureTestingModule({
        declarations: [HeptagonGridRowComponent, HeptagonComponent],
        imports: [SharedModule],
        providers: [{ provide: WindowToken, useFactory: windowProvider }],
      }).compileComponents();
    })
  );

  beforeEach(() => {
    fixture = TestBed.createComponent(HeptagonGridRowComponent);
    component = fixture.componentInstance;

    component.statuses = [
      { name: 'pod-1', status: 'ok' },
      { name: 'pod-2', status: 'ok' },
    ];
    component.edgeLength = 7;
    component.row = 1;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it(
    'should report when a heptagon is hovered',
    waitForAsync(() => {
      fixture.whenStable().then(() => {
        let got: HoverStatus;
        component.hoverState.subscribe((status: HoverStatus) => (got = status));

        component.updateHover(true, 1);

        const expected: HoverStatus = {
          row: component.row,
          col: 1,
          hovered: true,
        };
        expect(got).toEqual(expected);
      });
    })
  );
});
