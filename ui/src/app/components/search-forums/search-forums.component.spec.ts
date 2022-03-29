import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SearchForumsComponent } from './search-forums.component';

describe('SearchForumsComponent', () => {
  let component: SearchForumsComponent;
  let fixture: ComponentFixture<SearchForumsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SearchForumsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SearchForumsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
