import NamedCollection from './namedCollection'
import Note from '../note'
import groupBy from 'lodash.groupby'
import { today } from '@/util/dateHelper'

export default class GroupedCollection implements IterableIterator<NamedCollection> {
  groups: NamedCollection[]
  private pointer: number = 0
  
  constructor(notes: Note[], groupingFunc: Function, sortingFunc?: (a: any, b: any) => number) {
    const grouped = groupBy(notes, groupingFunc)

    let keys: string[]|number[] = Object.keys(grouped)
    
    if (sortingFunc) {
      keys = Object.keys(grouped).sort(sortingFunc)
    }

    this.groups = keys.map((key: string|number): NamedCollection => {
      return new NamedCollection(key, grouped[key])
    })
  }

  public next(): IteratorResult<NamedCollection> {
    if (this.pointer < this.groups.length) {
      return {
        done: false,
        value: this.groups[this.pointer++]
      }
    } else {
      this.pointer = 0
      return {
        done: true,
        value: null
      }
    }
  }

  [Symbol.iterator](): IterableIterator<NamedCollection> {
    return this
  }

  static createEmpty = () => {
    return new GroupedCollection([] as Note[], () => today)
  }
}