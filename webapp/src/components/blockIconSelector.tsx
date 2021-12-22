// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
import React, {useCallback} from 'react'

import {BlockIcons} from '../blockIcons'
import {Board} from '../blocks/board'
import {Card} from '../blocks/card'
import mutator from '../mutator'
import IconSelector from './iconSelector'
import EmojiPicker from '../widgets/emojiPicker'
import DeleteIcon from '../widgets/icons/delete'
import EmojiIcon from '../widgets/icons/emoji'
import Menu from '../widgets/menu'
import MenuWrapper from '../widgets/menuWrapper'

type Props = {
    block: Card
    size?: 's' | 'm' | 'l'
    readonly?: boolean
}

const BlockIconSelector = React.memo((props: Props) => {
    const {block, size} = props

    const onSelectEmoji = useCallback((emoji: string) => {
        mutator.changeBlockIcon(block.id, block.fields.icon, emoji)
        document.body.click()
    }, [block.id, block.fields.icon])
    const onAddRandomIcon = useCallback(() => mutator.changeBlockIcon(block.id, block.fields.icon, BlockIcons.shared.randomIcon()), [block.id, block.fields.icon])
    const onRemoveIcon = useCallback(() => mutator.changeBlockIcon(block.id, block.fields.icon, '', 'remove icon'), [block.id, block.fields.icon])

    if (!block.fields.icon) {
        return null
    }

    let className = `octo-icon size-${size || 'm'}`
    if (props.readonly) {
        className += ' readonly'
    }
    const iconElement = <div className={className}><span>{block.fields.icon}</span></div>

    return (
        <IconSelector
            readonly={props.readonly}
            iconElement={iconElement}
            onAddRandomIcon={onAddRandomIcon}
            onSelectEmoji={onSelectEmoji}
            onRemoveIcon={onRemoveIcon}
        />
    )
})

export default BlockIconSelector
