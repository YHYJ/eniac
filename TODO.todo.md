  - [ ] 考虑使用专门的 i18n 包代替 general/define_i18n.go (2024-05-29 16:24)
    - [ ] 需要将 i18n 文件打到包内，参考 skynet (2024-05-29 16:25)
  - [X] 完善输出格式 (2023-04-21 16:26)
  - [X] 添加生成示例配置文件功能 (2023-04-21 15:25)
  - [X] `get`子命令新增以下功能 (2023-04-20 14:35)
    - [X] 查询可用更新 (2023-04-20 14:36)
    - [X] 查询系统启动用时 (2023-04-20 14:36)
    - [X] 查询更新检测状态 (2023-04-20 14:36)
  - [X] 新增检测到配置文件错误时使用默认配置 (2023-05-12 13:34)
  - [X] 解决系统磁盘不是PCI设备的问题 (2023-10-10 09:47)
  - [X] 取消 'no such file' 报错 (2024-05-23 21:34)
  - [X] 先更新样式：使用 github.com/charmbracelet/lipgloss 替换 github.com/olekukonko/tablewriter 输出信息 (2024-05-30 14:49)
  - [X] 再更新布局，现在是所有信息一次输出，改成 Tab 样式的，参考 https://github.com/charmbracelet/bubbletea/tree/master/examples/tabs (2024-05-30 14:51)
  - [X] 检测系统是否能提供 update 信息 (2024-06-13 14:44)
  - [X] 从配置文件获取到的某一项的 items 长度为0时不显示该项 (2024-06-13 15:56)
  - [X] 完善构建约束 (2024-06-04 12:03)
  - [X] 现在官方库和 AUR 的可更新包列表是混杂在一个 cell 里的，想办法分开（横向排列可能出现过长的问题） (2024-12-19 13:33)
