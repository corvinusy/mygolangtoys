import QtQuick 2.0

Rectangle {
    width: 360
    height: 360
    color: "grey"

    Text {
        id: windowText
        anchors.centerIn: parent
        text: "Hello QML in Go!"
    }
}